// Initialize charts
let trendChart;
let categoryChart;

// Update charts with new data
async function updateCharts() {
    const token = localStorage.getItem('token');
    if (!token) return;

    try {
        // Fetch data for trend chart
        const trendResponse = await fetch(`${API_URL}/expenses/trend`, {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });
        const trendData = await trendResponse.json();
        updateTrendChart(trendData);

