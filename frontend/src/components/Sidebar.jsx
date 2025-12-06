import React from 'react';
import { 
  Clock, Star, FileText, Users, Folder, 
  Plus, Settings, Moon, ChevronRight 
} from 'lucide-react';

// Helper Component for individual Sidebar Items
const SidebarItem = ({ icon: Icon, label, level = 0, isActive = false, hasChildren = false }) => {
  const activeClasses = isActive
    ? 'bg-tint-amber/20 text-accent-amber font-semibold'
    : 'text-text-secondary hover:bg-white/60 hover:text-text-primary';

  // Indentation logic
  const paddingLeft = 16 + (level * 12);

  return (
    <div
      className={`group flex items-center py-1.5 pr-3 text-sm rounded-md cursor-pointer transition-all duration-200 mb-0.5 mx-2 ${activeClasses}`}
      style={{ paddingLeft: `${paddingLeft}px` }}
    >
      <div className="flex items-center flex-grow">
        {hasChildren ? (
            <ChevronRight className={`h-3 w-3 mr-2 text-text-secondary/50 ${isActive ? 'rotate-90' : ''}`} />
        ) : (
            <Icon className={`h-4 w-4 mr-3 opacity-80 group-hover:opacity-100 ${isActive ? 'text-accent-amber' : ''}`} />
        )}
        <span className="truncate">{label}</span>
      </div>
      {isActive && <div className="h-1.5 w-1.5 rounded-full bg-accent-amber ml-2" />}
    </div>
  );
};

const Sidebar = () => {
  const navigationGroups = [
    {
      label: "Navigation",
      items: [
        { icon: Clock, label: "Recent" },
        { icon: Star, label: "Starred" },
        { icon: FileText, label: "All Documents" },
      ],
    },
    {
      label: "Workspace",
      items: [
        { icon: Users, label: "Shared with me" },
        { icon: Folder, label: "Collections" },
      ],
    },
  ];

  const yourStuff = [
    { icon: FileText, label: "Project Phoenix", isActive: true },
    { icon: FileText, label: "Q4 Goals", level: 1 },
    { icon: FileText, label: "Marketing Sync", level: 1 },
    { icon: FileText, label: "Design System", level: 0 },
  ];

  return (
    <nav className="fixed top-[52px] left-0 h-[calc(100vh-52px)] w-[260px] border-r border-border-soft bg-bg-sidebar/95 backdrop-blur-md z-20 flex flex-col">
      
      {/* A. Vertical Accent Strip (Signature Identity) */}
      <div className="absolute top-0 left-0 bottom-0 w-1 bg-gradient-to-b from-tint-amber via-accent-amber to-hover-amber shadow-[1px_0_4px_rgba(217,119,6,0.3)]" />

      <div className="flex-grow overflow-y-auto py-4">
        {/* B. Workspace Block */}
        <div className="px-4 mb-6">
          <div className="p-3 bg-white border border-border-soft rounded-xl shadow-sm hover:shadow-md transition-shadow cursor-pointer flex items-center space-x-3 group">
            <div className="h-9 w-9 rounded-full bg-gradient-to-br from-stone-100 to-stone-200 border border-stone-200 flex items-center justify-center text-text-secondary font-semibold text-xs group-hover:border-accent-amber/30 transition-colors">
              AS
            </div>
            <div className="flex-col flex">
              <span className="font-semibold text-sm text-text-primary">Alok Singh</span>
              <span className="text-[11px] text-text-secondary uppercase tracking-wide font-medium">Free Workspace</span>
            </div>
          </div>
        </div>

        {/* C. Navigation Groups */}
        {navigationGroups.map((group, idx) => (
          <div key={idx} className="mb-6">
            <p className="px-5 mb-2 text-[11px] font-bold uppercase tracking-wider text-text-secondary/60">
              {group.label}
            </p>
            {group.items.map((item, i) => (
              <SidebarItem key={i} {...item} />
            ))}
          </div>
        ))}

        {/* D. Your Stuff Section */}
        <div className="mb-2">
            <div className="flex items-center justify-between px-5 mb-2 group cursor-pointer">
                <p className="text-[11px] font-bold uppercase tracking-wider text-text-secondary/60 group-hover:text-accent-amber transition-colors">
                Your Stuff
                </p>
                <Plus className="h-3 w-3 text-text-secondary opacity-0 group-hover:opacity-100 transition-opacity" />
            </div>
            
            {/* New Doc Button */}
            <div className="px-4 mb-3">
                <button className="w-full flex items-center justify-center py-2 text-sm font-medium rounded-lg text-white bg-accent-amber hover:bg-hover-amber transition-all shadow-sm hover:shadow active:scale-[0.98]">
                    <Plus className="h-4 w-4 mr-2" />
                    New Page
                </button>
            </div>

            {yourStuff.map((item, i) => (
              <SidebarItem key={i} {...item} />
            ))}
        </div>
      </div>

      {/* E. Footer */}
      <div className="p-4 border-t border-border-soft bg-bg-sidebar">
        <div className="flex items-center justify-between text-text-secondary">
          <button className="flex items-center space-x-2 px-2 py-1.5 rounded-md hover:bg-border-soft/50 transition-colors text-xs font-medium">
            <Settings className="h-4 w-4" />
            <span>Settings</span>
          </button>
          <button className="p-1.5 rounded-md hover:bg-border-soft/50 transition-colors">
             <Moon className="h-4 w-4" />
          </button>
        </div>
      </div>
    </nav>
  );
};

export default Sidebar;