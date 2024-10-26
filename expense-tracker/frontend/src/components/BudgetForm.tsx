import { useState, useEffect } from 'react';
import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Button,
  TextField,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  Grid,
  InputAdornment,
} from '@mui/material';
import { Budget } from '../types';
import { useAppContext } from '../context/AppContext';
import { budgetApi } from '../services/api';

