## Notes

- Why should we return 400 on GET/rate request? This is not a client error, because we don't pass any data. 
  It would be better return 503 in case of impossibility to get rate from external service.

- In case of really big user's database it would be better have separated requests for sending email and for obtaining results

  **Approaches:**

   1. User sends POST/sendEmails

   2. Server creates background task(s) and returns unied task id, like 

   ``` 
   http_response=202, {task_id:"1111-2222-3333-4444-5555"}
   ```

   3. User polls endpoint GET/task/{task_id} to get the status of the task, and updates frontend app with a detailed information about the task state as below:

    ```
    {status: 'in_progress', progress: 0.5, result: null} 
    ```
    or  
    
    ```
    {status: 'done', progress: 1.0, result: {success: 100, failed: 0}}
    ```

- For highloaded app consider cachig values in order to reduce number of requests to 3d party services, or throttling user's requests. Of cource if only we do not provide platform for HFT :)
