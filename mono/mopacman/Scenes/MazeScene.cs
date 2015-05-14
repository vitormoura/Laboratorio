using Microsoft.Xna.Framework;
using Microsoft.Xna.Framework.Graphics;
using Microsoft.Xna.Framework.Media;
using mopacman.Components;
using mopacman.Controllers;
using mopacman.Services;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace mopacman.Scenes
{
    class MazeScene : DrawableGameComponent
    {
        public Maze Maze { get; private set; }

        public KeyboardController Keyboard { get; private set; }

        public Texture2D Background { get; private set; }

        public Song Music { get; private set; }

        
        public MazeScene(MyGame g)
            : base(g)
        {
        }
                
        public override void Initialize()
        {
            MyGame game = this.Game as MyGame;

            this.Maze = MazeBuilder.GetDefaultFor(game.Content);

            this.PrepareMazeUI();

            Puckman p = new Puckman(game);
            p.CurrentLocation = this.Maze.GetStartSection();
            p.Initialize();

            this.Game.Components.Add(p);

            this.Keyboard = new KeyboardController(game, p);
            this.Keyboard.Initialize();
            
            //Ghost 1
            RegisterNewGhost("blinky.png", p, this.Maze[1, 4], this.Maze[5, 4]);
            
            //Ghost 2
            RegisterNewGhost("pinky.png", p, this.Maze[26, 22], this.Maze[29, 22]);

            //Ghost 3
            RegisterNewGhost("inky.png", p, this.Maze[26, 6], this.Maze[29, 6]);

            //Ghost 2
            RegisterNewGhost("clyde.png", p, this.Maze[1, 24], this.Maze[5, 24]);
                        
            base.Initialize();
        }

        private void PrepareMazeUI()
        {
            ///* Renderizando paredes como blocos
            foreach (var s in this.Maze)
            {
                if (s.Allowed)
                {
                    Block b = new Block(this.Game as MyGame, s);
                    b.SetPosition(new Point((int)(s.ID.X * b.Bounds.Width), (int)(s.ID.Y * b.Bounds.Height)));
                    b.Initialize();

                    this.Game.Components.Add(b);
                }
            }
            //*/
        }

        private void RegisterNewGhost(String ghostType, Puckman p, MazeSection r1, MazeSection r2)
        {
            Ghost g1 = new Ghost(this.Game as MyGame, ghostType);
            g1.Region = Tuple.Create(r1, r2);
            g1.CurrentLocation = this.Maze.GetGhostLairSection();
            g1.Initialize();

            this.Game.Components.Add(g1);

            GhostAIController iaCtrl1 = new GhostAIController(this.Game as MyGame, g1, p);
            iaCtrl1.Initialize();

            this.Game.Components.Add(iaCtrl1);
        }

        protected override void LoadContent()
        {
            this.Background = this.Game.Content.Load<Texture2D>("maze_template_1.png");
                        
            base.LoadContent();
        }

        public override void Draw(GameTime gameTime)
        {
            var sb = this.Game.Services.GetService<SpriteBatch>();
            sb.Draw(this.Background, position: Vector2.Zero);

            base.Draw(gameTime);
        }

        public override void Update(GameTime gameTime)
        {
            base.Update(gameTime);
        }
    }
}
