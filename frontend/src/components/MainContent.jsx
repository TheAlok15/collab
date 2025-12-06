import React from 'react';
import { FileText, ChevronRight } from 'lucide-react';

const MainContent = () => {
  return (
    <main className="flex-grow flex flex-col items-center justify-center h-full bg-bg-main relative overflow-hidden">
      {/* Background decoration: Subtle dots */}
      <div className="absolute inset-0 opacity-[0.4]" 
           style={{ backgroundImage: 'radial-gradient(#E7E5E4 1px, transparent 1px)', backgroundSize: '24px 24px' }}>
      </div>

      <div className="relative z-10 flex flex-col items-center max-w-md text-center p-8 animate-in fade-in zoom-in duration-500">
        
        {/* Large Empty State Icon */}
        <div className="w-24 h-24 bg-gradient-to-b from-bg-main to-bg-sidebar rounded-3xl border border-border-soft shadow-xl shadow-stone-200/50 flex items-center justify-center mb-8 rotate-3 transform hover:rotate-6 transition-transform duration-500">
           <FileText className="h-10 w-10 text-accent-amber" strokeWidth={1.5} />
        </div>

        <h1 className="text-3xl font-bold text-text-primary mb-3 tracking-tight">
          Welcome to Your Workspace
        </h1>
        
        <p className="text-text-secondary mb-8 leading-relaxed">
          Ready to get productive? Select an item from the sidebar or create a new document to get started.
        </p>

        <button className="group relative px-6 py-3 rounded-xl bg-accent-amber text-white font-medium shadow-lg shadow-accent-amber/25 hover:shadow-accent-amber/40 hover:bg-hover-amber transition-all active:translate-y-0.5">
          <span className="flex items-center">
             Create New Document
             <ChevronRight className="h-4 w-4 ml-2 opacity-70 group-hover:translate-x-1 transition-transform" />
          </span>
        </button>

        <div className="mt-12 flex space-x-6 text-xs text-text-secondary/60 font-medium">
            <span>Press <kbd className="font-sans px-1.5 py-0.5 rounded border border-border-soft bg-white/50">⌘ K</kbd> to search</span>
            <span><kbd className="font-sans px-1.5 py-0.5 rounded border border-border-soft bg-white/50">N</kbd> for new page</span>
        </div>
      </div>
    </main>
  );
};

export default MainContent;