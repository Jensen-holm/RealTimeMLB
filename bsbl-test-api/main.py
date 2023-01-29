from flask import Flask, request, jsonify
from pybaseball import statcast


app = Flask(__name__)


@app.route("/", methods=["POST"])
def main():
    return
