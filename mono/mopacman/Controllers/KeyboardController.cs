using Microsoft.Xna.Framework;
using Microsoft.Xna.Framework.Input;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace mopacman.Controllers
{
    class KeyboardController : GameComponent
    {
        private IControllable player;
        private EnumDirections lastDirection;
        
        public KeyboardController(MyGame g, IControllable player)
            : base(g)
        {
            this.player = player;
        }

        public override void Update(GameTime gameTime)
        {
            KeyboardState state = Keyboard.GetState();
            
            if (state.IsKeyDown(Keys.Left))
            {
                this.lastDirection = EnumDirections.West;
            }
            else if (state.IsKeyDown(Keys.Right))
            {
                this.lastDirection = EnumDirections.East;
            }
            else if (state.IsKeyDown(Keys.Up))
            {
                this.lastDirection = EnumDirections.North;
            }
            else if (state.IsKeyDown(Keys.Down))
            {
                this.lastDirection = EnumDirections.South;
            }
            else
            {
                return;
            }

            this.player.GoTo(this.lastDirection);
        }
    }
}
