
import json
import tornado.ioloop

from tornado import web
from tornadoes import ESConnection


es = ESConnection('local.elasticsearch.com', '9200')

class MainHandler(tornado.web.RequestHandler):

    @web.gen.coroutine
    def get(self):
        res = yield es.search(index='sitemap_g1', body={"query": {"match_all": {}}})
        data = json.loads(res.body.decode('utf-8'))
        self.render('templates/index.xml',
                    sitemaps=data['hits']['hits'])

def make_app():
    return web.Application([
        (r"/", MainHandler),
    ])

app = make_app()

if __name__ == "__main__":
    app.listen(8888)
    tornado.ioloop.IOLoop.current().start()
