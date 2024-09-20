// DOM Elements
const authSection = document.getElementById('auth-section');
const dashboardSection = document.getElementById('dashboard-section');
const loginForm = document.getElementById('login-form');
const registerForm = document.getElementById('register-form');
const tabBtns = document.querySelectorAll('.tab-btn');
const authTabs = document.querySelector('.auth-tabs');
const togglePasswordBtns = document.querySelectorAll('.toggle-password');

// Password strength configuration
const passwordStrengthConfig = {
    minLength: 8,
    patterns: {
        hasNumber: /\d/,
        hasLetter: /[a-zA-Z]/,
        hasSpecial: /[!@#$%^&*(),.?":{}|<>]/,
        hasUppercase: /[A-Z]/,
        hasLowercase: /[a-z]/
    }
};

// Password validation rules
const passwordRules = {
    minLength: 8,
    maxLength: 32,
    requireUppercase: true,
    requireLowercase: true,
    requireNumber: true,
    requireSpecial: true,
    allowedSpecial: "!@#$%^&*(),.?\":{}|<>",
    recommendedLength: 12,
    uniqueCharsMin: 6,
    noCommonPatterns: true,
    commonPatterns: [
        '123', '234', '345', '456', '567', '678', '789', '987', '876', '765', '654', '543', '432', '321',
        'password', 'qwerty', 'asdfgh', 'zxcvbn'
    ]
};

function calculatePasswordStrength(password) {
    let score = 0;
    const maxScore = 100;
    
    // Length score (30%)
    const lengthScore = Math.min((password.length - passwordRules.minLength) / 
        (passwordRules.recommendedLength - passwordRules.minLength), 1) * 30;
    score += lengthScore;
    
    // Character variety score (40%)
    const hasUpper = /[A-Z]/.test(password);
    const hasLower = /[a-z]/.test(password);
    const hasNumber = /\d/.test(password);
    const hasSpecial = new RegExp(`[${passwordRules.allowedSpecial}]`).test(password);
    
    score += hasUpper ? 10 : 0;
    score += hasLower ? 10 : 0;
    score += hasNumber ? 10 : 0;
