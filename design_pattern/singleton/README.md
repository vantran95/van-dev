**1️⃣ What is the Singleton Pattern?**

The **Singleton Pattern** is a **creational design pattern** that ensures **only one instance** of a struct is created and provides a **global access point** to it.

**✅ When to Use the Singleton Pattern?**

•	When you need **only one instance** of an object throughout the application.

•	When you want to **share a single resource** (e.g., a database connection, logger, or configuration).

**2️⃣ Implementing Singleton in Go**

Unlike other languages, Go does not have **class-based objects**, so we implement the Singleton pattern using **structs, sync.Once, and package-level variables**.

**🔹 Basic Singleton Implementation**

```go
package main

import (
	"fmt"
	"sync"
)

// Singleton struct (e.g., a database connection)
type Singleton struct {
	data string
}

var instance *Singleton
var once sync.Once

// GetInstance ensures only one instance is created
func GetInstance() *Singleton {
	once.Do(func() {
		fmt.Println("Creating single instance now...")
		instance = &Singleton{data: "Singleton Instance"}
	})
	return instance
}

func main() {
	// First call: instance is created
	s1 := GetInstance()
	fmt.Println(s1.data)

	// Second call: returns the same instance
	s2 := GetInstance()
	fmt.Println(s2.data)

	// Check if both instances are the same
	fmt.Println("Are instances the same?", s1 == s2)
}
```

**🔹 Output**

```
Creating single instance now...
Singleton Instance
Singleton Instance
Are instances the same? true
```

**🔹 How It Works?**

1.	We use **a package-level variable (instance)** to store the single instance.

2.	We use **sync.Once** to ensure the instance is created only once, even in concurrent environments.

3.	Calling GetInstance() multiple times **returns the same instance**.

**3️⃣ Use Cases of the Singleton Pattern**

✅ **1. Database Connection Pooling**

→ Avoid multiple connections by using a **single shared connection**.

✅ **2. Logging Service**

→ Maintain a **single log manager** to avoid race conditions.

✅ **3. Configuration Manager**

→ Store global **app configurations** in a single instance.

✅ **4. Caching**

→ Maintain a **global cache** (e.g., Redis client).

**4️⃣ Thread-Safe Singleton with Lazy Initialization**

For high-performance applications, use **sync.Once** to ensure **thread safety**.

**🔍 Expected Behavior**

1.	Only **one instance** of Logger is created.

2.	Multiple goroutines log messages **without conflicts**.

3.	Logs are written to app.log in a **thread-safe manner**.

**🚀 Real-World Applications of Singleton Pattern**

| **Use Case** | **Description** |
| --- | --- |
| **Logging System** | Prevent multiple loggers from writing to different files. |
| **Database Connection Pool** | Ensure a **single connection pool** is shared across the application. |
| **Configuration Manager** | Centralized **global settings** that are accessed throughout the app. |
| **Cache (e.g., Redis Client)** | A single instance of **cache client** prevents redundant connections. |
| **Authentication Service** | Maintain a single **JWT token manager** instance for security. |

**🏆 Key Takeaways**

✅ **Ensures a single instance** of an object exists.

✅ **Thread-safe** using sync.Once.

✅ **Global access** to a shared resource.

✅ Prevents **redundant object creation**, saving memory and resources.

**🚀 Conclusion**

The **Singleton Pattern** is a powerful design pattern that ensures only **one instance** of an object exists. In this case study, it helped us implement a **thread-safe logging system** for concurrent access in a **scalable web application**. 🎯