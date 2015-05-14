using Microsoft.Xna.Framework;
using mopacman.Controllers;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace mopacman.Components
{
    class Ghost : Player
    {
        public enum States
        {
            Waiting,
            Chase,
            Scatter,
            Frightened
        }

        public States State
        {
            get { return this.behavior.State; }
        }

        public Tuple<MazeSection,MazeSection> Region { get; set; }
               
        public Ghost(MyGame g)
            : base(g, "ghost.png", new Rectangle(0, 0, Constants.DEFAULT_BLOCK_WIDTH, Constants.DEFAULT_BLOCK_WIDTH))
        {
            this.behavior = new GhostBehavior();
        }

        public override void Update(GameTime gameTime)
        {
            this.behavior.Update(gameTime);
        }

        private GhostBehavior behavior;
    }
}
