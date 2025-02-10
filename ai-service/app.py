from flask import Flask, jsonify
from flask_cors import CORS
import nltk
from nltk.sentiment.vader import SentimentIntensityAnalyzer

nltk.download('vader_lexicon')

app = Flask(__name__)
CORS(app)

sia = SentimentIntensityAnalyzer()

@app.route('/', methods=['GET'])
def home():
  return jsonify({"message": "Python AI Service is running!"})

@app.route('/api/sentiment', methods=['POST'])
def sentiment_analysis():
  test_text = "I love building stuff!"
  sentiment = sia.polarity_scores(test_text)
  return jsonify({"text": test_text, "sentiment": sentiment})

if __name__ == '__main__':
  app.run(port=5001, debug=True)