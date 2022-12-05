FROM python:3.11 AS requirements-stage

WORKDIR /tmp

COPY api/Pipfile* /tmp/

RUN pip install --no-cache-dir pipenv

RUN pipenv requirements > requirements.txt

FROM python:3.11

WORKDIR /app

COPY --from=requirements-stage /tmp/requirements.txt /app/requirements.txt

RUN pip install --no-cache-dir --upgrade pip && \
    pip install --no-cache-dir --upgrade -r /app/requirements.txt

RUN groupadd --gid 1000 fastapi && \
    useradd --uid 1000 --gid 1000 -m fastapi

COPY --chown=fastapi:fastapi api app/api

USER fastapi

CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "80"]

EXPOSE 80
