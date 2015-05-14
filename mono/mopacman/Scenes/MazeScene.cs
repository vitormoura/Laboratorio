using Microsoft.Xna.Framework;
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

        public MazeScene(MyGame g)
            : base(g)
        {
        }

        public override void Initialize()
        {
            MyGame game = this.Game as MyGame;

            this.Maze = MazeBuilder.GetDefaultFor(game.Content);

            foreach (var s in this.Maze)
            {
                if (!s.Allowed)
                {
                    Block b = new Block(this.Game as MyGame);
                    b.SetPosition(new Point((int)(s.ID.X * b.Bounds.Width), (int)(s.ID.Y * b.Bounds.Height)));
                    b.Initialize();

                    this.Game.Components.Add(b);
                }
            }

            Puckman p = new Puckman(game);
            p.CurrentLocation = this.Maze.GetStartSection();
            p.Initialize();

            this.Game.Components.Add(p);

            this.Keyboard = new KeyboardController(game, p);
            this.Keyboard.Initialize();
            
            //Ghost 1
            RegisterNewGhost(p, this.Maze[1, 4], this.Maze[5, 4]);
            
            //Ghost 2
            RegisterNewGhost(p, this.Maze[23, 20], this.Maze[29, 20]);
                        
            base.Initialize();
        }

        private void RegisterNewGhost(Puckman p, MazeSection r1, MazeSection r2)
        {
            Ghost g1 = new Ghost(this.Game as MyGame);
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
            base.LoadContent();
        }

        public override void Draw(GameTime gameTime)
        {
            base.Draw(gameTime);
        }

        public override void Update(GameTime gameTime)
        {
            base.Update(gameTime);
        }
    }
}
