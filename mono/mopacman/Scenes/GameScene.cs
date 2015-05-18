using Microsoft.Xna.Framework;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace mopacman.Scenes
{
    abstract class GameScene : DrawableGameComponent
    {
        public  List<GameComponent> Components { get; protected set; }
                
        public GameScene(MyGame g)
            : base(g)
        {
            this.Components = new List<GameComponent>();

            this.End();
        }

        public override void Update(GameTime gameTime)
        {
            if (this.Enabled)
            {
                foreach (GameComponent c in this.Components)
                {
                    c.Update(gameTime);
                }

                base.Update(gameTime);
            }
        }

        public override void Draw(GameTime gameTime)
        {
            if (this.Enabled && this.Visible)
            {
                foreach (GameComponent c in this.Components)
                {
                    DrawableGameComponent d = c as DrawableGameComponent;

                    if (d != null)
                        d.Draw(gameTime);
                }

                base.Draw(gameTime);
            }
        }

        public void End()
        {
            this.Visible = false;
            this.Enabled = false;
        }

        public void Begin()
        {
            this.Visible = true;
            this.Enabled = true;
        }
    }
}
