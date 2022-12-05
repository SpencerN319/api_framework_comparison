from dotenv import dotenv_values
from fastapi import FastAPI
from pymongo import MongoClient
from .greeting import greeting_routes
from .vulnerabilities import vulnerability_routes

config = dotenv_values(".env")

app = FastAPI()
app.include_router(greeting_routes.router)
app.include_router(vulnerability_routes.router)


@app.on_event("startup")
def init():
    app.mongodb_client = MongoClient(config["MONGO_URL"])
    app.database = app.mongodb_client[config["DB_NAME"]]
    print("Connected to the MongoDB database!")


@app.get("/")
async def root():
    return {"message": "Hello World"}


@app.on_event("shutdown")
def tear_down():
    app.mongodb_client.close()
