sudo kubectl exec -it frontend-deployment-5f65f54684-zpmsr -- curl -X POST backend-service:5000/submit   -H "Content-Type: application/json"   -d '{"name": "Mayur Satish Khadde", "email": "mayur.khadde22@vit.edu", "feedback": "123"}'   



k exec -it mongo-client -- sh

mongosh "mongodb://mongodb-service:27017"

db.forms.find({name: "Mayur Satish Khadde"})


sudo kubectl exec -it frontend-deployment-5f65f54684-64n6x -- curl -X POST backend-service:5000/submit   -H "Content-Type: application/json"   -d '{"name": "Mayur Satish Kh
adde", "email": "mayur.khadde22@vit.edu", "feedback": "123"}'



kubectl run curl-pod --image=appropriate/curl --restart=Never -- sleep 3600


kubectl exec -it curl-pod -- sh

k exec -it curl-pod -- curl -X POST backend-service:5000/submit   -H "Content-Type: application/json"   -d '{"name": "Mayur Satish Khadde", "email": "mayur.khadde22@vit.edu", "feedback": "123"}'   


kubectl rollout restart deployment <deployment-name>


kubectl config set-context --current --namespace=project