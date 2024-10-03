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

        // Fetch data for category distribution
        const categoryResponse = await fetch(`${API_URL}/expenses/by-category`, {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });
        const categoryData = await categoryResponse.json();
        updateCategoryChart(categoryData);
    } catch (error) {
        console.error('Failed to update charts:', error);
    }
}

// Update expense trend chart
function updateTrendChart(data) {
    const ctx = document.getElementById('trend-chart').getContext('2d');
    
    // Destroy existing chart if it exists
    if (trendChart) {
        trendChart.destroy();
    }

    trendChart = new Chart(ctx, {
        type: 'line',
        data: {
            labels: data.labels,
            datasets: [{
                label: 'Daily Expenses',
                data: data.values,
                borderColor: '#2563eb',
                backgroundColor: 'rgba(37, 99, 235, 0.1)',
                tension: 0.4,
                fill: true
            }]
        },
        options: {
            responsive: true,
            maintainAspectRatio: false,
