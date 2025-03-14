import React, { useEffect } from 'react';
import './ToastMessage.css';

type ToastMessageProps = {
  message: string;
  type: 'success' | 'error';
  onClose: () => void;
};

const ToastMessage: React.FC<ToastMessageProps> = ({ message, type, onClose }) => {
  useEffect(() => {
    const timer = setTimeout(() => {
      onClose();
    }, 3000); // Auto-hide after 3 seconds

    return () => clearTimeout(timer);
  }, [onClose]);

  return (
    <div className={`toast ${type}`}>
      {message}
    </div>
  );
};

export default ToastMessage;
