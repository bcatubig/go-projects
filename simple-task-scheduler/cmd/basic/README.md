
Add tasks to scheduler with various unit lengths

## Output

```
2020-08-22T20:26:20.869-0500 [INFO]  simple-task-scheduler.scheduler-1: worker added: scheduler_id=35aba330-6a00-4535-a610-54ec1f5d78f1 worker_id=4902bf47-a9c9-4f75-9709-efe8f2a93774
2020-08-22T20:26:20.870-0500 [INFO]  simple-task-scheduler.scheduler-1: worker added: scheduler_id=35aba330-6a00-4535-a610-54ec1f5d78f1 worker_id=39000ec3-65d2-4dcf-bc8b-4ca810e83be6
2020-08-22T20:26:20.870-0500 [INFO]  simple-task-scheduler.worker-1: task added: task_id=97b811cf-f2a1-4745-b6b4-94a2fe200d81 available_work_units=2
2020-08-22T20:26:20.870-0500 [INFO]  simple-task-scheduler.scheduler-1: task added to worker: task_id=97b811cf-f2a1-4745-b6b4-94a2fe200d81 worker_id=4902bf47-a9c9-4f75-9709-efe8f2a93774
2020-08-22T20:26:20.870-0500 [INFO]  simple-task-scheduler.scheduler-1: not able to add task to worker: worker_id=4902bf47-a9c9-4f75-9709-efe8f2a93774 error="not able to add task. require 3 work units, 2 available"
2020-08-22T20:26:20.870-0500 [INFO]  simple-task-scheduler.worker-2: task added: task_id=1395dc71-d2c4-4669-8dec-49af521ccfc0 available_work_units=0
2020-08-22T20:26:20.870-0500 [INFO]  simple-task-scheduler.scheduler-1: task added to worker: task_id=1395dc71-d2c4-4669-8dec-49af521ccfc0 worker_id=39000ec3-65d2-4dcf-bc8b-4ca810e83be6
2020-08-22T20:26:20.870-0500 [INFO]  simple-task-scheduler.worker-1: task added: task_id=cccbed3f-5b8f-45c9-9845-ac9507f2f255 available_work_units=0
2020-08-22T20:26:20.870-0500 [INFO]  simple-task-scheduler.scheduler-1: task added to worker: task_id=cccbed3f-5b8f-45c9-9845-ac9507f2f255 worker_id=4902bf47-a9c9-4f75-9709-efe8f2a93774
2020-08-22T20:26:20.870-0500 [INFO]  simple-task-scheduler.scheduler-1: not able to add task to worker: worker_id=4902bf47-a9c9-4f75-9709-efe8f2a93774 error="not able to add task. require 5 work units, 0 available"
2020-08-22T20:26:20.870-0500 [INFO]  simple-task-scheduler.scheduler-1: not able to add task to worker: worker_id=39000ec3-65d2-4dcf-bc8b-4ca810e83be6 error="not able to add task. require 5 work units, 0 available"
2020-08-22T20:26:20.870-0500 [ERROR] simple-task-scheduler.scheduler-1: no workers available to schedule task: task_id=93b138be-0473-4026-ad94-6e5b55b4cac5
2020-08-22T20:26:20.870-0500 [ERROR] simple-task-scheduler: failed to add task: error="no workers available"
```
