from flask import Flask,render_template,jsonify,request
import json

app= Flask(__name__)
val =10
@app.route('/')
def runner():
    return("hello world")

@app.route('/',methods=["GET"])
def runner2():
    return(jsonify({'message':"bye bye"}))

@app.route("/gettut",methods=["GET"])
def getter():
    return jsonify({"message":f'yess , you got the message {val}'})

@app.route('/posttut',methods=["POST"])
def poster():
    global val
    value = request.get_json()
    value = dict(value)
    val=value['value']
    return jsonify({'message':f'you updated the value to {val}'})

if __name__ == '__main__':
    app.run(debug=True)