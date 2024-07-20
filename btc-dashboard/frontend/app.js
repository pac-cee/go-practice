// Fetch and display the current BTC price
async function fetchPrice() {
  try {
    const res = await fetch('http://localhost:8088/price');
    const data = await res.json();
    document.getElementById('price').textContent = `BTC/USDT: $${data.price.toLocaleString(undefined, {maximumFractionDigits: 2})}`;
  } catch (err) {
    document.getElementById('price').textContent = 'Error loading price';
  }
}

// Fetch and display candlestick chart
async function fetchCandles() {
  try {
    const res = await fetch('http://localhost:8088/candles');
    const candles = await res.json();
    renderCandlestickChart(candles);
  } catch (err) {
    const ctx = document.getElementById('candlestickChart').getContext('2d');
