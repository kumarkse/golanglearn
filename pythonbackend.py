from flask import Flask,render_template,jsonify
import json

app= Flask(__name__)

@app.route('/')
def runner():
    return("hello world")

@app.route('/',methods=["GET"])
def runner2():
    return(jsonify({'message':"bye bye"}))

if __name__ == '__main__':
    app.run(debug=True)