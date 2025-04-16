FROM scratch
COPY httpservertimeout /
# RUN chmod +x /httpservertimeout
EXPOSE 8080
CMD ["/httpservertimeout"]