### DDD

![ddd](https://miro.medium.com/max/600/1*VqlXUp6QvmijjBFdQp2ssg.jpeg)

DDD has 4 layers in the architecture:
1. **Interface**: This layer responsibles for the interaction with user, whether software presents information or recieves information from user.
2. **Application**: This is a thin layer between interface and domain, it could call domain services to serve the application purposes.
3. **Domain**: The heart of the software, this layer holds domain logic and business knowledge.
4. **Infrastructure**: A supporting layer for the other layers. This layer contains supporting libraries or external services like database or UI supporting library.
