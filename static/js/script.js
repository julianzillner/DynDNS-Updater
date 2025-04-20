document.addEventListener("DOMContentLoaded", function () {
    // /health abfragen
    fetch("/health")
        .then(response => response.text())
        .then(data => {
            document.getElementById("health").textContent = data;
        })
        .catch(error => {
            document.getElementById("health").textContent = "Error beim Laden des Status.";
            console.error("/health Error:", error);
        });

    // /currentIP abfragen
    fetch("/currentIP")
        .then(response => response.text())
        .then(data => {
            document.getElementById("currentIP").textContent = data;
        })
        .catch(error => {
            document.getElementById("currentIP").textContent = "Error beim Laden der IP.";
            console.error("/currentIP Error:", error);
        });


        // /currentIP abfragen
    fetch("/interval/sec")
        .then(response => response.text())
        .then(data => {
            document.getElementById("interval").textContent = data;
        })
        .catch(error => {
            document.getElementById("interval").textContent = "Error beim Laden der IP.";
            console.error("/interval Error:", error);
        });
});
