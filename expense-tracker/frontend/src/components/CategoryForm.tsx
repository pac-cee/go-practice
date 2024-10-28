import { useState, useEffect } from 'react';
import {
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Button,
  TextField,
  Box,
  Chip,
  IconButton,
} from '@mui/material';
import { Add as AddIcon, Close as CloseIcon } from '@mui/icons-material';
import { Category } from '../types';
import { useAppContext } from '../context/AppContext';
import { categoryApi } from '../services/api';

interface CategoryFormProps {
  open: boolean;
