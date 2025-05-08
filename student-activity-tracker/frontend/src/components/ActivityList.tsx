import React from 'react';

interface Activity {
  id: string;
  title: string;
  description: string;
  duration: number;
  status: string;
  startTime?: string;
  endTime?: string;
}

interface ActivityListProps {
  activities: Activity[];
  onStart: (id: string) => void;
  onComplete: (id: string) => void;
  onDelete: (id: string) => void;
}

