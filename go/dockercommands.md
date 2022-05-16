<!-- docker build -f /home/dom/mercari-build-training-2022/go/dockerfile -t build2022/app:latest /home/dom/mercari-build-training-2022/go/
 -->
docker build --tag build2022/app .
docker run --publish 9000:9000 build2022/app
