from flask import Flask, render_template

from elasticsearch import Elasticsearch

es = Elasticsearch(['http://local.elasticsearch.com:80'])

app = Flask(__name__)

@app.route('/')
def index():
    res = es.search(index="sitemap_g1", body={"query": {"match_all": {}}})
    return render_template("index.xml",
                           sitemaps=res['hits']['hits'],
                           headers={'Content-Type': 'application/xml'})


if __name__ == "__main__":
    app.run()
