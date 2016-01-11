GoRelic
=======

NewRelic middleware for gin-gonic framework.

## Usage

```go
import(
	"github.com/gin-gonic/gin"
	"github.com/brandfolder/gin-gorelic"
)

func main(){
	g := gin.Default()

	handler, err := gorelic.InitNewrelicAgent("YOUR_NEWRELIC_LICENSE_KEY", "YOUR_APP_NAME", true, nil)
	if err == nil {
		router.Use(handler)
	} else {
		fmt.Printf("Something went wrong initialising NewRelic %v\n", err)
	}

	g.Run()
}
```

## Defining custom metrics
Here's a general definition. For more details take a look into the examples.


```go
import(
	"github.com/gin-gonic/gin"
	"github.com/brandfolder/gin-gorelic"
)

type M struct {
	name        string
	units       string
	counter     float64
}

func (_m *M) GetName() string  { return "CustomCounters/" + _m.name }
func (_m *M) GetUnits() string { return _m.units }
func (_m *M) GetValue() (float64, error) {
	//Method body...
}

func main(){
	g := gin.Default()
	
	m1 := &M{name: "custom_metric_1", units: "unit", counter: 1}
	m2 := &M{name: "custom_metric_2", units: "unit", counter: 2}
	customMetrics := []gorelic.Metric{m1, m2}

	handler, err := gorelic.InitNewrelicAgent("YOUR_NEWRELIC_LICENSE_KEY", "YOUR_APPLICATION_NAME", true, customMetrics)
	if err == nil {
		router.Use(handler)
	} else {
		fmt.Printf("Something went wrong initialising NewRelic %v\n", err)
	}

	g.Run()
}
```


## Authors

* [Jason Waldrip](http://github.com/jwaldrip)
* [Yuriy Vasiyarov](http://github.com/yvasiyarov) [original martini gorelic]

## Contributions

* [Ronald Arias](https://github.com/ronald05arias)