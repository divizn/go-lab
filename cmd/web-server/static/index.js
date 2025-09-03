document.addEventListener("DOMContentLoaded", () => {
  let btn = document.getElementById("clickMe");
  let output = document.getElementById("output");

  btn.addEventListener("click", async () => {
    try {
      let res = await fetch("/api");
      if (!res.ok) throw new Error(`HTTP ${res.status}`);
      let data = await res.json();
      output.textContent = JSON.stringify(data, null, 2);
    } catch (err) {
      output.textContent = `Error: ${err.message}`;
    }
  });
});
