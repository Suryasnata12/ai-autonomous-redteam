async function analyze() {
  const target = document.getElementById("target").value;
  const output = document.getElementById("output");

  if (!target) {
    output.textContent = "Please enter a domain";
    return;
  }

  output.textContent = "Running analysis...";

  try {
    const res = await fetch("/api/analyze", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({ target })
    });

    if (!res.ok) {
      const err = await res.json();
      throw new Error(err.error || "Request failed");
    }

    const data = await res.json();
    output.textContent = JSON.stringify(data, null, 2);

  } catch (err) {
    output.textContent = "Error: " + err.message;
  }
}
