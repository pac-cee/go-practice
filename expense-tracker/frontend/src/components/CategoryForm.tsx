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
  onClose: () => void;
  category?: Category;
}

const CategoryForm = ({ open, onClose, category }: CategoryFormProps) => {
  const { dispatch } = useAppContext();
  const [name, setName] = useState('');
  const [newSubcategory, setNewSubcategory] = useState('');
  const [subcategories, setSubcategories] = useState<string[]>([]);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (category) {
      setName(category.name);
      setSubcategories(category.subcategories);
    }
  }, [category]);

  const handleAddSubcategory = () => {
