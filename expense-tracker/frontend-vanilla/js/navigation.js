// Navigation module
const navigation = {
    // Current active page
    currentPage: window.location.pathname.split('/').pop().replace('.html', ''),

    // Initialize navigation
    init() {
        this.setupNavigation();
        this.highlightCurrentPage();
        this.setupLogout();
    },

    // Set up navigation event listeners
    setupNavigation() {
        const navLinks = document.querySelectorAll('.nav-link');
        navLinks.forEach(link => {
            link.addEventListener('click', (e) => {
                if (link.classList.contains('logout')) return;
                
