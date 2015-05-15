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
        private EnumDirections nextDirection;
        
        public KeyboardController(MyGame g, IControllable player)
            : base(g)
        {
            this.player = player;
            this.player.ReadyToMove += player_ReadyToMove;
        }
                
        public override void Update(GameTime gameTime)
        {
            KeyboardState state = Keyboard.GetState();
            
            if (state.IsKeyDown(Keys.Left))
            {
                this.nextDirection = EnumDirections.West;
            }
            else if (state.IsKeyDown(Keys.Right))
            {
                this.nextDirection = EnumDirections.East;
            }
            else if (state.IsKeyDown(Keys.Up))
            {
                this.nextDirection = EnumDirections.North;
            }
            else if (state.IsKeyDown(Keys.Down))
            {
                this.nextDirection = EnumDirections.South;
            }
        }

        private void player_ReadyToMove(object sender, EventArgs e)
        {
            var nextSection = this.player.CurrentLocation.Get(this.nextDirection);

            if (nextSection != null && nextSection.Allowed)
            {
                this.lastDirection = nextDirection;
                this.player.GoTo(nextDirection);
            }
            else
            {
                this.player.GoTo(this.lastDirection);
            }
        }
    }
}
