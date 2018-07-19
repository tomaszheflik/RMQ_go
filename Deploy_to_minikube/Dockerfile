FROM rabbitmq:latest
ENV RABBITMQ_USE_LONGNAME=true
ENV AUTOCLUSTER_LOG_LEVEL=debug
ENV AUTOCLUSTER_CLEANUP=true
ENV CLEANUP_INTERVAL=60
ENV CLEANUP_WARN_ONLY=false
ENV AUTOCLUSTER_TYPE=k8s
#ENV LANG=en_US.UTF-8
ADD plugins/*.ez /usr/lib/rabbitmq/lib/rabbitmq_server-3.7.7/plugins/
RUN rabbitmq-plugins enable --offline autocluster 
RUN rabbitmq-plugins enable --offline rabbitmq_management
