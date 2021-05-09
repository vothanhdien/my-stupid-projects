package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"time"
)

var data = map[string]string{
	"DSC09716.JPG": "1zhBlV7fw_a0j5esiLVisx4umOpgYwDyl",
	"DSC09672.JPG": "1uLMiwebxFleH_l8KjjSBveRpveLzui0_",
	"DSC09665.JPG": "1omzq9LkM53ZN-_Vn5z6pLfYqny5jemmt",
	"DSC09687.JPG": "1v7WsN_TO8bgRI6VSd1uGslr3QuXR2HdJ",
	"DSC09700.JPG": "1JVoiEYQ8nBapsBNV9gP66dPEGuqc4-OY",
	"DSC09682.JPG": "194ZkqFiukWP9EJruj1EC3bo7sKXjstiI",
	"DSC09719.JPG": "1M6QGPjT4tkduv4ywz2mK7TuS-WcwH5VY",
	"DSC09720.JPG": "1_5wZnDNTDyS8LwFOtYs9AbPodp1pfs8G",
	"DSC09692.JPG": "1FpExPK2VuZRfiMiAcRV95S9Z1w2YWA_X",
	"DSC09690.JPG": "1Bu0-9zamCP1qCi0loZ-HD9HcSa3pPY11",
	"DSC09751.JPG": "1Dk_rK-bt9dhWuZdn7vgAePdSmp5GzpRv",
	"DSC09754.JPG": "19-effKun384QSfmSUixvAG7ANDMpWWWj",
	"DSC09763.JPG": "17iW2-YhZxSNXtrShIQYOV3zqLkhewxQc",
	"DSC09773.JPG": "170Xhiv8Bj6TzH4ufVT3qBZZ_HXKzdEBr",
	"DSC09753.JPG": "1ZunnFNJkNEdNjgREyXmK8G8LUGENWrTW",
	"DSC09771.JPG": "1ZRXpeHdiXweRRIEkS1n80TKW-Unl0rTV",
	"DSC09731.JPG": "13joSJBXkMK9pErxjJHqbk910OrPMhXas",
	"DSC09774.JPG": "1T-7P8pItAwKp_N5BE8OSFrlivaBWl7RJ",
	"DSC09756.JPG": "1T7H4qfcDaH-UfnZ8j6OYOZtu8H54gIcf",
	"DSC09786.JPG": "1BJLVvfNcz8vt9PB9bJKtcqUz4JSM11M8",
}

var tp = &http.Transport{
	MaxIdleConnsPerHost: 21,
}

func main() {
	err := make(chan error, 1)

	for name, id := range data {
		go func(id, name string, c chan error) {
			result, size, duration, err := downloadImage(id, name)
			if err != nil {
				c <- err
			}
			fmt.Println(id, name, result, size, duration)
		}(id, name, err)
	}
	fmt.Printf("num goroutine %v\n", runtime.NumGoroutine())
	e := <-err
	panic(e)
}

func downloadImage(id string, name string) (string, int64, time.Duration, error) {
	start := time.Now()
	f, err := os.Create("photo/" + name)
	if err != nil {
		return "", 0, time.Since(start), err
	}
	defer f.Close()

	c := http.Client{
		Transport: tp,
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	url := fmt.Sprintf("https://drive.google.com/uc?export=download&id=%v", id)
	r, err := c.Get(url)
	if err != nil {
		fmt.Printf("Error while downloading %q: %v", url, err)
		return "", 0, time.Since(start), err
	}
	defer r.Body.Close()

	n, err := io.Copy(f, r.Body)
	if err != nil {
		return "copy error", 0, time.Since(start), err
	}

	return "success", n, time.Since(start), nil
}
