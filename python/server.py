# server.py
from concurrent import futures
import grpc
import analysis_pb2
import analysis_pb2_grpc
import pandas as pd
from sklearn.linear_model import LinearRegression

class AnalysisService(analysis_pb2_grpc.AnalysisServiceServicer):
    def AnalyzeData(self, request, context):
        # Convert gRPC MarketDataRequest to pandas DataFrame
        data = [{'time': md.time, 'symbol': md.symbol, 'price': md.price, 'volume': md.volume} for md in request.data]
        df = pd.DataFrame(data)

        # Perform data analysis
        df['moving_average'] = df['price'].rolling(window=20).mean().fillna(0)

        # Machine learning model
        X = df[['volume']].values
        y = df['price'].values
        model = LinearRegression().fit(X, y)
        predictions = model.predict(X).tolist()

        # Prepare the response
        response = analysis_pb2.AnalysisResponse(
            predictions=predictions,
            moving_averages=df['moving_average'].tolist()
        )
        return response

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    analysis_pb2_grpc.add_AnalysisServiceServicer_to_server(AnalysisService(), server)
    server.add_insecure_port('[::]:42069')
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    serve()
