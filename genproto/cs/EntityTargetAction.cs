// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: EntityTargetAction.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Deviant {

  /// <summary>Holder for reflection information generated from EntityTargetAction.proto</summary>
  public static partial class EntityTargetActionReflection {

    #region Descriptor
    /// <summary>File descriptor for EntityTargetAction.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static EntityTargetActionReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "ChhFbnRpdHlUYXJnZXRBY3Rpb24ucHJvdG8SB0RldmlhbnQaC1RpbGVzLnBy",
            "b3RvIj8KEkVudGl0eVRhcmdldEFjdGlvbhIKCgJpZBgBIAEoCRIdCgV0aWxl",
            "cxgCIAEoCzIOLkRldmlhbnQuVGlsZXNCCVoHZGV2aWFudGIGcHJvdG8z"));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Deviant.TilesReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Deviant.EntityTargetAction), global::Deviant.EntityTargetAction.Parser, new[]{ "Id", "Tiles" }, null, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  public sealed partial class EntityTargetAction : pb::IMessage<EntityTargetAction> {
    private static readonly pb::MessageParser<EntityTargetAction> _parser = new pb::MessageParser<EntityTargetAction>(() => new EntityTargetAction());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<EntityTargetAction> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Deviant.EntityTargetActionReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public EntityTargetAction() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public EntityTargetAction(EntityTargetAction other) : this() {
      id_ = other.id_;
      tiles_ = other.tiles_ != null ? other.tiles_.Clone() : null;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public EntityTargetAction Clone() {
      return new EntityTargetAction(this);
    }

    /// <summary>Field number for the "id" field.</summary>
    public const int IdFieldNumber = 1;
    private string id_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string Id {
      get { return id_; }
      set {
        id_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "tiles" field.</summary>
    public const int TilesFieldNumber = 2;
    private global::Deviant.Tiles tiles_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Deviant.Tiles Tiles {
      get { return tiles_; }
      set {
        tiles_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as EntityTargetAction);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(EntityTargetAction other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (Id != other.Id) return false;
      if (!object.Equals(Tiles, other.Tiles)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (Id.Length != 0) hash ^= Id.GetHashCode();
      if (tiles_ != null) hash ^= Tiles.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (Id.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(Id);
      }
      if (tiles_ != null) {
        output.WriteRawTag(18);
        output.WriteMessage(Tiles);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (Id.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Id);
      }
      if (tiles_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Tiles);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(EntityTargetAction other) {
      if (other == null) {
        return;
      }
      if (other.Id.Length != 0) {
        Id = other.Id;
      }
      if (other.tiles_ != null) {
        if (tiles_ == null) {
          Tiles = new global::Deviant.Tiles();
        }
        Tiles.MergeFrom(other.Tiles);
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            Id = input.ReadString();
            break;
          }
          case 18: {
            if (tiles_ == null) {
              Tiles = new global::Deviant.Tiles();
            }
            input.ReadMessage(Tiles);
            break;
          }
        }
      }
    }

  }

  #endregion

}

#endregion Designer generated code
