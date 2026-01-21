fetch("/report")
  .then(r => r.json())
  .then(data => {
    document.getElementById("output").textContent =
      JSON.stringify(data, null, 2);
  });
