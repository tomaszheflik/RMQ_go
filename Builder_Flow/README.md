#### Builder Flow

[App: Generate build job] --> |MQ: init_job| --> [App: Visualise] --> |MQ: build| --> [App: Builder app] --> |MQ: status|

|MQ: status| --> [App: update_status] -ch-> [App: Visualise]



