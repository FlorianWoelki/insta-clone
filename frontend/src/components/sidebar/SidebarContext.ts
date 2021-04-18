import { Context, createContext } from 'react';

type SidebarContextType = {
  activeItem: number;
};

const SidebarContext: Context<SidebarContextType> = createContext({
  activeItem: 0,
});

export default SidebarContext;
