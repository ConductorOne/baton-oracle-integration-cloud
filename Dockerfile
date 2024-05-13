FROM gcr.io/distroless/static-debian11:nonroot
ENTRYPOINT ["/baton-oracle-integration-cloud"]
COPY baton-oracle-integration-cloud /