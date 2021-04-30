import React from 'react';

export type ButtonType = 'primary' | 'secondary';
type ButtonProps = {
  type?: ButtonType;
};

const Button: React.FC<ButtonProps> = ({
  children,
  type = 'primary',
}): JSX.Element => {
  const colorClasses = {
    primary:
      'text-white bg-gradient-to-r from-pink-600 to-yellow-500 hover:from-pink-800 hover:to-yellow-700',
    secondary:
      'text-pink-600 border border-pink-600 hover:text-white bg-gradient-to-r hover:from-pink-600 hover:to-yellow-500',
  };

  return (
    <button
      type="button"
      className={`flex items-center px-4 py-2 space-x-1 font-medium rounded-lg focus:outline-none ${colorClasses[type]}`}>
      {children}
    </button>
  );
};

Button.defaultProps = {
  type: 'primary',
};

export default Button;
