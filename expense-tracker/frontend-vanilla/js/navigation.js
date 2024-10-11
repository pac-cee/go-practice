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
                
                e.preventDefault();
                const page = link.getAttribute('href');
                this.navigateTo(page);
            });
        });
    },

    // Navigate to a page
    navigateTo(page) {
        window.location.href = page;
    },

    // Highlight current page in navigation
    highlightCurrentPage() {
        const navLinks = document.querySelectorAll('.nav-link');
        navLinks.forEach(link => {
            const linkPage = link.getAttribute('href').split('/').pop().replace('.html', '');
            link.classList.toggle('active', linkPage === this.currentPage);
        });
