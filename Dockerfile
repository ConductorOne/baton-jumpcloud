FROM gcr.io/distroless/static-debian11:nonroot
ENTRYPOINT ["/baton-jumpcloud"]
COPY baton-jumpcloud /