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

            KeyboardController kbCtrl = new KeyboardController(game, p);
            kbCtrl.Initialize();

            this.Game.Components.Add(kbCtrl);



            Ghost g = new Ghost(this.Game as MyGame);
            g.CurrentLocation = this.Maze.GetGhostLairSection();
            g.Initialize();
                        
            this.Game.Components.Add(g);
            
            BasicIAController iaCtrl = new BasicIAController(this.Game as MyGame, g, p);
            iaCtrl.Initialize();

            this.Game.Components.Add(iaCtrl);

            base.Initialize();
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
