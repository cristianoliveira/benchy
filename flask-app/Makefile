run:
	gunicorn --workers 4 --bind 0.0.0.0:5000 app:app

bench:
	wrk -c 10 -d 60 -t 10 "http://localhost:8800"
