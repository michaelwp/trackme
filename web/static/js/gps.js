if (navigator.geolocation) {
    navigator.geolocation.getCurrentPosition(
        function(position) {
            console.log("Latitude:", position.coords.latitude);
            console.log("Longitude:", position.coords.longitude);
            // window.location.replace("https://www.google.com/maps/place/Sapporo,+Hokkaido,+Japan/@42.9853655,140.9183319,10z/data=!3m1!4b1!4m6!3m5!1s0x5f0ad4755a973633:0x33937e9d4687bad5!8m2!3d43.0617713!4d141.3544507!16zL20vMGdwNWw2?entry=ttu&g_ep=EgoyMDI1MDcyOS4wIKXMDSoASAFQAw%3D%3D");
        },
        function(error) {
            console.error("Error getting location:", error.message);
        }
    );
} else {
    console.error("Geolocation is not supported by this browser.");
}