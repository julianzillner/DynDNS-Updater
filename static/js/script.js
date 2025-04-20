function showSection(id) {
  document.querySelectorAll('.content-section').forEach((el) => el.classList.remove('active'));
  document.getElementById(id).classList.add('active');
}

// Optional: direkt den Bereich anzeigen, wenn # in URL vorhanden
window.addEventListener('DOMContentLoaded', () => {
  const hash = window.location.hash.replace('#', '');
  if (hash) {
    showSection(hash);
  }
});

document.addEventListener('DOMContentLoaded', function () {
  // /health abfragen
  fetch('/health')
    .then((response) => response.text())
    .then((data) => {
      document.getElementById('health').textContent = data;
    })
    .catch((error) => {
      document.getElementById('health').textContent = 'Error beim Laden des Status.';
      console.error('/health Error:', error);
    });

  // /currentIP abfragen
  fetch('/currentIP')
    .then((response) => response.text())
    .then((data) => {
      document.getElementById('currentIP').textContent = data;
    })
    .catch((error) => {
      document.getElementById('currentIP').textContent = 'Error beim Laden der IP.';
      console.error('/currentIP Error:', error);
    });

  // /currentIP abfragen
  fetch('/interval/min')
    .then((response) => response.text())
    .then((data) => {
      document.getElementById('interval').textContent = data;
    })
    .catch((error) => {
      document.getElementById('interval').textContent = 'Error beim Laden der IP.';
      console.error('/interval Error:', error);
    });
});
