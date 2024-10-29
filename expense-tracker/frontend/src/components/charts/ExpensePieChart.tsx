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
