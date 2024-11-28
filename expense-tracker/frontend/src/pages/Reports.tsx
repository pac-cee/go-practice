import { Grid, Paper, Typography, Box, Select, MenuItem, FormControl, InputLabel } from '@mui/material';
import { useState } from 'react';

const Reports = () => {
  const [timeRange, setTimeRange] = useState('month');

  return (
    <Grid container spacing={3}>
      <Grid item xs={12} display="flex" justifyContent="space-between" alignItems="center">
        <Typography variant="h4">Reports</Typography>
        <FormControl sx={{ minWidth: 200 }}>
          <InputLabel>Time Range</InputLabel>
          <Select
            value={timeRange}
            label="Time Range"
            onChange={(e) => setTimeRange(e.target.value)}
          >
            <MenuItem value="week">Last Week</MenuItem>
            <MenuItem value="month">Last Month</MenuItem>
