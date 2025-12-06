import React from 'react';
import { Search, Bell, User } from 'lucide-react';

const Header = () => {
  return (
    <header className="flex justify-between items-center h-[52px] px-4 border-b border-border-soft bg-bg-main text-text-primary fixed top-0 left-0 right-0 z-30 shadow-sm">
      {/* Left: App Logo */}
      <div className="flex items-center w-[260px] flex-shrink-0">
        <div className="flex items-center justify-center h-8 w-8 rounded-lg bg-gradient-to-br from-tint-amber to-accent-amber text-white font-bold shadow-md">
          C
        </div>
      </div>

      {/* Center: Document Title */}
      <div className="flex-grow flex justify-center">
        <div className="px-4 py-1.5 rounded-md hover:bg-border-soft/30 cursor-pointer transition-colors">
            <span className="text-sm font-medium text-text-secondary">No document selected</span>
        </div>
      </div>

      {/* Right: Global Actions */}
      <div className="flex items-center space-x-3">
        <button className="p-2 rounded-md hover:bg-border-soft/50 text-text-secondary hover:text-text-primary transition-colors">
            <Search className="h-4 w-4" />
        </button>
        <button className="p-2 rounded-md hover:bg-border-soft/50 text-text-secondary hover:text-text-primary transition-colors relative">
            <Bell className="h-4 w-4" />
            <span className="absolute top-2 right-2 h-2 w-2 bg-accent-amber rounded-full border-2 border-bg-main"></span>
        </button>
        <button className="ml-1 h-8 w-8 rounded-full bg-border-soft flex items-center justify-center text-text-secondary hover:ring-2 ring-border-soft transition-all">
            <User className="h-4 w-4" />
        </button>
      </div>
    </header>
  );
};

export default Header;