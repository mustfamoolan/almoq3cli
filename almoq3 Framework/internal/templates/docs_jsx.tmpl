import { useState } from 'react'
import { Link } from 'react-router-dom'

const NAV = [
  { id: 'introduction', label: '👋 Introduction' },
  { id: 'installation', label: '📦 Installation' },
  { id: 'quick-start', label: '🚀 Quick Start' },
  { id: 'project-structure', label: '📁 Project Structure' },
  { id: 'commands', label: '🛠️ CLI Commands' },
  { id: 'authentication', label: '🔐 Authentication' },
  { id: 'database', label: '🗄️ Database & Migrations' },
  { id: 'frontend', label: '🎨 Frontend Engine' },
  { id: 'upgrade', label: '⬆️ Upgrade System' },
  { id: 'deployment', label: '🐳 Deployment' },
  { id: 'roadmap', label: '🗺️ Roadmap' },
]

const Code = ({ children, block }) => (
  <code style={ {
    backgroundColor: '#0d1a2d',
    color: '#00BFFF',
    padding: block ? '1rem 1.25rem' : '0.15rem 0.4rem',
    borderRadius: '6px',
    fontFamily: 'monospace',
    fontSize: block ? '0.875rem' : '0.85em',
    display: block ? 'block' : 'inline',
    whiteSpace: block ? 'pre' : 'normal',
    overflowX: block ? 'auto' : 'visible',
    lineHeight: block ? '1.7' : 'inherit',
    border: '1px solid #1e3a5f',
  } }>{children}</code>
)

const Section = ({ id, title, children }) => (
  <section id={id} style={ { marginBottom: '3rem', scrollMarginTop: '80px' } }>
    <h2 style={ {
      fontSize: '1.5rem',
      fontWeight: '700',
      color: '#0D2338',
      borderBottom: '2px solid #00BFFF',
      paddingBottom: '0.5rem',
      marginBottom: '1.25rem',
    } }>{title}</h2>
    {children}
  </section>
)

const P = ({ children }) => (
  <p style={ { color: '#475569', lineHeight: '1.75', marginBottom: '0.875rem', fontSize: '0.9375rem' } }>{children}</p>
)

const CommandRow = ({ cmd, desc }) => (
  <tr>
    <td style={ { padding: '0.625rem 0.75rem', borderBottom: '1px solid #e2e8f0', width: '42%' } }>
      <Code>{cmd}</Code>
    </td>
    <td style={ { padding: '0.625rem 0.75rem', borderBottom: '1px solid #e2e8f0', color: '#475569', fontSize: '0.875rem' } }>{desc}</td>
  </tr>
)

export default function Docs() {
  const [activeSection, setActiveSection] = useState('introduction')

  const scrollTo = (id) => {
    setActiveSection(id)
    document.getElementById(id)?.scrollIntoView({ behavior: 'smooth' })
  }

  return (
    <div style={ { minHeight: '100vh', backgroundColor: '#f8fafc', fontFamily: "'Instrument Sans', ui-sans-serif, system-ui, sans-serif" } }>

      {/* Top Nav */}
      <nav style={ {
        position: 'fixed', top: 0, left: 0, right: 0, zIndex: 50,
        backgroundColor: '#0D2338', padding: '0 2rem',
        display: 'flex', alignItems: 'center', justifyContent: 'space-between',
        height: '60px', boxShadow: '0 2px 8px rgba(0,0,0,0.2)',
      } }>
        <Link to="/" style={ { display: 'flex', alignItems: 'center', gap: '0.625rem', textDecoration: 'none' } }>
          <svg width="28" height="28" viewBox="0 0 40 40" fill="none">
            <rect width="40" height="40" rx="8" fill="#0D2338"/>
            <path d="M8 28L20 8L32 28H25L20 19L15 28H8Z" fill="#00BFFF"/>
          </svg>
          <span style={ { color: '#ffffff', fontWeight: '600', fontSize: '1.1rem' } }>almoq3</span>
          <span style={ { color: '#00BFFF', fontSize: '0.75rem', backgroundColor: 'rgba(0,191,255,0.15)', padding: '0.125rem 0.5rem', borderRadius: '4px', border: '1px solid rgba(0,191,255,0.3)' } }>docs</span>
        </Link>
        <div style={ { display: 'flex', gap: '1rem', alignItems: 'center' } }>
          <Link to="/" style={ { color: '#94a3b8', textDecoration: 'none', fontSize: '0.875rem' } }>← Back to Home</Link>
          <a href="https://github.com/mustfamoolan/almoq3cli" target="_blank" style={ { color: '#00BFFF', textDecoration: 'none', fontSize: '0.875rem' } }>GitHub →</a>
        </div>
      </nav>

      <div style={ { display: 'flex', paddingTop: '60px', maxWidth: '1200px', margin: '0 auto' } }>

        {/* Sidebar */}
        <aside style={ {
          width: '260px', flexShrink: 0, position: 'sticky', top: '60px',
          height: 'calc(100vh - 60px)', overflowY: 'auto',
          padding: '2rem 1rem', borderRight: '1px solid #e2e8f0',
          backgroundColor: '#ffffff',
        } }>
          <p style={ { fontSize: '0.75rem', fontWeight: '600', color: '#94a3b8', textTransform: 'uppercase', letterSpacing: '0.08em', marginBottom: '0.75rem' } }>Contents</p>
          {NAV.map(item => (
            <button key={item.id} onClick={() => scrollTo(item.id)} style={ {
              display: 'block', width: '100%', textAlign: 'left',
              padding: '0.5rem 0.75rem', borderRadius: '6px', border: 'none', cursor: 'pointer',
              fontSize: '0.875rem', marginBottom: '2px',
              backgroundColor: activeSection === item.id ? '#eff6ff' : 'transparent',
              color: activeSection === item.id ? '#0D2338' : '#64748b',
              fontWeight: activeSection === item.id ? '600' : '400',
              borderLeft: activeSection === item.id ? '3px solid #00BFFF' : '3px solid transparent',
              transition: 'all 0.15s',
            } }>{item.label}</button>
          ))}
        </aside>

        {/* Main Content */}
        <main style={ { flex: 1, padding: '2.5rem 3rem', maxWidth: '800px' } }>

          <Section id="introduction" title="Introduction">
            <P>
              <strong>almoq3</strong> (المعقد) is a full-stack enterprise framework that brings the developer experience
              of <strong>Laravel</strong> to the raw performance of <strong>Go</strong>. It scaffolds production-ready
              applications with a single command, eliminating the configuration overhead that typically plagues Go projects.
            </P>
            <P>
              Think of it as <em>Laravel for Go</em> — but with React and shadcn/ui built-in from day one.
            </P>
            <div style={ { backgroundColor: '#eff6ff', border: '1px solid #bfdbfe', borderLeft: '4px solid #00BFFF', borderRadius: '6px', padding: '1rem 1.25rem', margin: '1rem 0' } }>
              <strong style={ { color: '#0D2338' } }>💡 Core Philosophy:</strong>
              <span style={ { color: '#475569', marginLeft: '0.5rem' } }>Zero boilerplate. Maximum productivity. Production-ready from day one.</span>
            </div>
          </Section>

          <Section id="installation" title="Installation">
            <P>almoq3 requires <strong>Go 1.21+</strong> and <strong>Node.js 18+</strong>.</P>
            <Code block>npm install -g almoq3-cli</Code>
            <P style={ { marginTop: '1rem' } }>Verify the installation:</P>
            <Code block>almoq3 --version{'\n'}# almoq3 version 2.1.1</Code>
          </Section>

          <Section id="quick-start" title="Quick Start">
            <Code block>{`# Create a new full-stack project
almoq3 new my-awesome-app

# Navigate into the project
cd my-awesome-app

# Start the server
go run main.go

# Open your browser
# http://localhost:3000`}</Code>
            <P>That's it. Your full-stack app — Go backend + React frontend — is live.</P>
          </Section>

          <Section id="project-structure" title="Project Structure">
            <Code block>{`my-app/
├── app/
│   ├── controllers/      # HTTP handlers (RESTful CRUD)
│   ├── models/           # GORM models
│   ├── services/         # Business logic layer
│   ├── middleware/        # Custom middleware
│   └── requests/         # Form validation
├── bootstrap/            # App initialization
├── config/               # Typed config (Viper)
├── database/
│   ├── migrations/       # SQL migrations
│   └── seeders/          # Database seeders
├── routes/
│   ├── api.go            # API routes
│   └── web.go            # Web routes + React serve
├── frontend/             # React + Vite + shadcn/ui
│   └── src/
│       ├── App.jsx
│       ├── Docs.jsx
│       └── UIShowcase.jsx
├── public/               # Built assets (auto-generated)
├── Dockerfile
├── docker-compose.yml
└── almoq3.json           # Version tracking`}</Code>
          </Section>

          <Section id="commands" title="CLI Commands">
            <div style={ { overflowX: 'auto' } }>
              <table style={ { width: '100%', borderCollapse: 'collapse', fontSize: '0.875rem' } }>
                <thead>
                  <tr style={ { backgroundColor: '#0D2338' } }>
                    <th style={ { padding: '0.75rem', textAlign: 'left', color: '#00BFFF', fontWeight: '600' } }>Command</th>
                    <th style={ { padding: '0.75rem', textAlign: 'left', color: '#ffffff', fontWeight: '600' } }>Description</th>
                  </tr>
                </thead>
                <tbody>
                  <CommandRow cmd="almoq3 new <name>" desc="Scaffold a complete full-stack project" />
                  <CommandRow cmd="almoq3 run" desc="Start the development server" />
                  <CommandRow cmd="almoq3 upgrade" desc="Safely upgrade framework files" />
                  <CommandRow cmd="almoq3 make:controller <Name>" desc="Generate a RESTful controller" />
                  <CommandRow cmd="almoq3 make:model <Name>" desc="Generate a GORM model" />
                  <CommandRow cmd="almoq3 make:service <Name>" desc="Generate a service layer class" />
                  <CommandRow cmd="almoq3 make:middleware <Name>" desc="Generate middleware" />
                  <CommandRow cmd="almoq3 make:request <Name>" desc="Generate form request validator" />
                  <CommandRow cmd="almoq3 make:migration <name>" desc="Generate a timestamped SQL migration" />
                  <CommandRow cmd="almoq3 migrate" desc="Run pending migrations" />
                  <CommandRow cmd="almoq3 migrate:rollback" desc="Roll back the last migration batch" />
                  <CommandRow cmd="almoq3 make:seeder <Name>" desc="Generate a database seeder" />
                  <CommandRow cmd="almoq3 db:seed" desc="Execute database seeders" />
                  <CommandRow cmd="almoq3 make:auth" desc="Scaffold complete JWT authentication" />
                  <CommandRow cmd="almoq3 make:job <Name>" desc="Generate a background job" />
                  <CommandRow cmd="almoq3 make:test <Name>" desc="Generate a test suite file" />
                  <CommandRow cmd="almoq3 key:generate" desc="Generate a secure APP_KEY" />
                  <CommandRow cmd="almoq3 self-update" desc="Update the CLI to the latest version" />
                </tbody>
              </table>
            </div>
          </Section>

          <Section id="authentication" title="Authentication (JWT)">
            <P>Generate a complete JWT authentication system with one command:</P>
            <Code block>almoq3 make:auth</Code>
            <P>This generates 4 production-ready files:</P>
            <ul style={ { color: '#475569', lineHeight: '2', paddingLeft: '1.25rem', marginBottom: '1rem' } }>
              <li><Code>app/controllers/AuthController.go</Code> — Register, Login, Logout, Me</li>
              <li><Code>app/models/User.go</Code> — User model with bcrypt hashing</li>
              <li><Code>app/middleware/JwtMiddleware.go</Code> — Route protection</li>
              <li><Code>database/migrations/xxxx_create_users_table.sql</Code></li>
            </ul>
            <P>Available endpoints after <Code>almoq3 migrate</Code>:</P>
            <Code block>{`POST /api/auth/register
POST /api/auth/login
GET  /api/auth/me       (protected)
POST /api/auth/logout   (protected)`}</Code>
          </Section>

          <Section id="database" title="Database & Migrations">
            <Code block>{`# Create a migration
almoq3 make:migration create_products_table

# Run migrations
almoq3 migrate

# Rollback
almoq3 migrate:rollback

# Seed the database
almoq3 db:seed`}</Code>
            <P>Supports: <strong>PostgreSQL</strong>, <strong>MySQL</strong>, <strong>SQLite</strong></P>
          </Section>

          <Section id="frontend" title="Frontend Engine">
            <P>Every almoq3 project ships with a fully configured React frontend:</P>
            <ul style={ { color: '#475569', lineHeight: '2', paddingLeft: '1.25rem', marginBottom: '1rem' } }>
              <li>⚡ <strong>Vite</strong> — Lightning-fast HMR</li>
              <li>🎨 <strong>Tailwind CSS</strong> — Utility-first styling</li>
              <li>🧩 <strong>shadcn/ui</strong> — Production-quality components</li>
              <li>🔲 <strong>Radix UI</strong> — Accessible primitives</li>
              <li>🔀 <strong>React Router</strong> — Client-side navigation</li>
            </ul>
            <P>The Go server automatically serves the built React assets and handles client-side routing fallback.</P>
            <Code block>{`# For development (hot reload)
cd frontend && npm run dev

# For production (served by Go)
cd frontend && npm run build
go run main.go`}</Code>
          </Section>

          <Section id="upgrade" title="Upgrade System">
            <P>almoq3 has a smart upgrade system that updates framework files without touching your code:</P>
            <Code block>almoq3 upgrade</Code>
            <div style={ { display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '1rem', margin: '1rem 0' } }>
              <div style={ { padding: '1rem', backgroundColor: '#f0fdf4', border: '1px solid #bbf7d0', borderRadius: '6px' } }>
                <p style={ { color: '#166534', fontWeight: '600', marginBottom: '0.5rem' } }>✅ Gets Updated</p>
                <ul style={ { color: '#475569', fontSize: '0.875rem', lineHeight: '1.8', paddingLeft: '1rem' } }>
                  <li>frontend/src/App.jsx</li>
                  <li>frontend/src/Docs.jsx</li>
                  <li>frontend/vite.config.js</li>
                  <li>routes/web.go</li>
                </ul>
              </div>
              <div style={ { padding: '1rem', backgroundColor: '#fef2f2', border: '1px solid #fecaca', borderRadius: '6px' } }>
                <p style={ { color: '#991b1b', fontWeight: '600', marginBottom: '0.5rem' } }>🛡️ Never Touched</p>
                <ul style={ { color: '#475569', fontSize: '0.875rem', lineHeight: '1.8', paddingLeft: '1rem' } }>
                  <li>app/controllers/</li>
                  <li>app/models/</li>
                  <li>routes/api.go</li>
                  <li>.env</li>
                </ul>
              </div>
            </div>
          </Section>

          <Section id="deployment" title="Deployment">
            <P>Every project includes Docker configuration out of the box:</P>
            <Code block>{`# Start full stack (App + PostgreSQL + Redis)
docker-compose up -d

# Or build manually
go build -o app main.go
./app`}</Code>
          </Section>

          <Section id="roadmap" title="Roadmap">
            {[
              { done: true, text: 'Core scaffolding engine' },
              { done: true, text: 'GORM + multi-database support' },
              { done: true, text: 'JWT authentication generator' },
              { done: true, text: 'React + shadcn/ui frontend integration' },
              { done: true, text: 'Smart project upgrade system' },
              { done: true, text: 'Background jobs (Redis)' },
              { done: true, text: 'Docker + production deployment' },
              { done: true, text: 'Update notifications in terminal' },
              { done: true, text: 'Built-in documentation page' },
              { done: false, text: 'almoq3 make:ui — scaffold shadcn components' },
              { done: false, text: 'WebSocket support' },
              { done: false, text: 'GraphQL generator' },
              { done: false, text: 'Official documentation website' },
              { done: false, text: 'VS Code extension' },
            ].map((item, i) => (
              <div key={i} style={ { display: 'flex', alignItems: 'center', gap: '0.75rem', padding: '0.5rem 0', borderBottom: '1px solid #f1f5f9' } }>
                <span style={ { fontSize: '1rem' } }>{item.done ? '✅' : '⏳'}</span>
                <span style={ { color: item.done ? '#0D2338' : '#94a3b8', fontSize: '0.9375rem', textDecoration: item.done ? 'none' : 'none' } }>{item.text}</span>
              </div>
            ))}
          </Section>

          <div style={ { textAlign: 'center', padding: '2rem 0', borderTop: '1px solid #e2e8f0', color: '#94a3b8', fontSize: '0.875rem' } }>
            Built with ❤️ in Iraq 🇮🇶 · <a href="https://github.com/mustfamoolan/almoq3cli" target="_blank" style={ { color: '#00BFFF', textDecoration: 'none' } }>GitHub</a>
          </div>
        </main>
      </div>
    </div>
  )
}
