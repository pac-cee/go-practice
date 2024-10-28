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
    if (newSubcategory.trim() && !subcategories.includes(newSubcategory.trim())) {
      setSubcategories([...subcategories, newSubcategory.trim()]);
      setNewSubcategory('');
    }
  };

  const handleDeleteSubcategory = (subcategory: string) => {
    setSubcategories(subcategories.filter((sub) => sub !== subcategory));
  };

  const handleSubmit = async () => {
    if (!name.trim()) return;

    try {
      setLoading(true);
      const data = {
        name: name.trim(),
        subcategories,
      };

      if (category) {
        const response = await categoryApi.update(category.id, data);
        dispatch({ type: 'UPDATE_CATEGORY', payload: response.data });
      } else {
        const response = await categoryApi.create(data);
        dispatch({ type: 'ADD_CATEGORY', payload: response.data });
      }
      onClose();
    } catch (error) {
      console.error('Error saving category:', error);
      dispatch({ type: 'SET_ERROR', payload: 'Error saving category' });
    } finally {
      setLoading(false);
    }
  };

  const handleClose = () => {
    setName('');
