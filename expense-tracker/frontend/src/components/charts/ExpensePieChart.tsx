import React from 'react';
import { Pie } from 'react-chartjs-2';
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js';
import { CategoryExpenseSummary } from '../../types';
import { Box, useTheme } from '@mui/material';

ChartJS.register(ArcElement, Tooltip, Legend);

interface ExpensePieChartProps {
  data: CategoryExpenseSummary[];
}

const ExpensePieChart: React.FC<ExpensePieChartProps> = ({ data }) => {
  const theme = useTheme();

  const chartData = {
    labels: data.map((item) => item.category),
    datasets: [
      {
        data: data.map((item) => item.amount),
        backgroundColor: [
          theme.palette.primary.main,
          theme.palette.secondary.main,
          theme.palette.error.main,
          theme.palette.warning.main,
          theme.palette.info.main,
          theme.palette.success.main,
          // Add more colors as needed
        ],
        borderColor: theme.palette.background.paper,
        borderWidth: 1,
      },
    ],
  };

  const options = {
    responsive: true,
    maintainAspectRatio: false,
